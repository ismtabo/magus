import { resolve } from "node:path";

import { remark } from "remark";
import remarkGfm from "remark-gfm";
import { gt } from "semver";
import { read, write } from "to-vfile";
import { dedent } from "ts-dedent";
import { find } from "unist-util-find";
import { findAfter } from "unist-util-find-after";
import { findAllAfter } from "unist-util-find-all-after";
import findAllBetween from "unist-util-find-all-between";
import { convert } from "unist-util-is";
import { remove } from "unist-util-remove";
import { visit } from "unist-util-visit";

import { isVersionSection, parseVersionHeader } from "./utils/unist.mjs";

export default async function main({ logger, args, options }) {
  logger.info(`Updating CHANGELOG.md to version ${args.version} (${args.date})`);
  const version = args.version;
  const date = args.date;
  const changelog = await read(resolve(process.cwd(), options.changelog));
  const newChangelog = await remark()
    .use(remarkGfm)
    .use(remarkCleanUnreleasedEmptySections, logger)
    .use(remarkReplaceUnreleased, version, date, logger)
    .process(changelog);
  newChangelog.value = newChangelog.value.replaceAll("\\[", "[");

  await write(newChangelog);
  logger.info("Done");
}

const isUnreleasedSection = convert(
  (node) =>
    node.type === "heading" &&
    node.depth === 2 &&
    node.children[0] &&
    node.children[0].value.includes("[Unreleased]"),
);

const isChangeSection = convert(
  (node) =>
    node.type === "heading" &&
    node.depth === 3 &&
    node.children[0] &&
    ["Added", "Changed", "Deprecated", "Removed", "Fixed", "Security"].includes(
      node.children[0].value,
    ),
);

function remarkCleanUnreleasedEmptySections(logger) {
  return function transformer(tree, file) {
    const unreleasedSection = find(tree, isUnreleasedSection);
    if (!unreleasedSection) {
      logger.error("Could not find Unreleased section");
      file.fail("Could not find Unreleased section", {
        source: "remark-clean-unreleased-empty-sections",
      });
    }
    const versionSection = findAfter(tree, unreleasedSection, isVersionSection);
    const changesSubsections = versionSection
      ? findAllBetween(tree, unreleasedSection, versionSection, isChangeSection)
      : findAllAfter(tree, unreleasedSection, isChangeSection);
    const emptyChangesSubsections = changesSubsections.filter((subsection) =>
      isEmptySection(tree, subsection),
    );
    if (!emptyChangesSubsections.length) {
      logger.debug("No empty changes subsections found");
      return;
    }
    logger.info(`Removing ${emptyChangesSubsections.length} empty changes subsections`);
    const isEmptyChangesSubSection = convert((node) => emptyChangesSubsections.includes(node));
    remove(tree, isEmptyChangesSubSection);
  };
}

const unreleasedSectionContent = dedent`
## [Unreleased]

### Added

### Changed

### Deprecated

### Removed

### Fixed

### Security
`;

function generateVersionHeader(version, date) {
  return `[${version}] - (${date.format("YYYY-MM-DD")})`;
}

function isEmptySection(tree, node) {
  const sibling = findAfter(tree, node);
  return !sibling || (sibling.type === "heading" && sibling.depth <= node.depth);
}

function remarkReplaceUnreleased(version, date, logger) {
  const newTree = remark().parse(unreleasedSectionContent);
  return function transformer(tree, file) {
    const versionHeader = generateVersionHeader(version, date);
    visit(tree, isUnreleasedSection, (node, index, parent) => {
      if (isEmptySection(tree, node)) {
        file.fail("Unreleased section is empty", {
          source: "remark-replace-unreleased",
        });
      }
      const sibling = findAfter(tree, node, isVersionSection);
      if (sibling) {
        const previousVersion = parseVersionHeader(sibling.children[0].value);
        if (previousVersion) {
          if (!gt(version, previousVersion.version)) {
            file.fail(`Version ${version} should be greater than ${previousVersion.version}`, {
              source: "remark-replace-unreleased",
            });
          }
        }
      }
      logger.info(`Replacing ${node.children[0].value} with ${version}`);
      if (index && parent) {
        parent.children.splice(index, 1, ...newTree.children, node);
      }
      node.children[0].value = versionHeader;
      return [visit.SKIP, index + newTree.children.length + 1];
    });
    return tree;
  };
}
