import { resolve } from "node:path";

import { remark } from "remark";
import { read } from "to-vfile";
import { find } from "unist-util-find";
import { findAfter } from "unist-util-find-after";
import { findAllAfter } from "unist-util-find-all-after";
import findAllBetween from "unist-util-find-all-between";
import { convert } from "unist-util-is";
import { reporter } from "vfile-reporter";

import { isVersionSection, parseVersionHeader } from "./utils/unist.mjs";

const isVersionSectionForSection = function (version) {
  return convert((node) => {
    if (!isVersionSection(node)) return false;
    const sectionVersion = parseVersionHeader(node.children[0].value);
    return sectionVersion && sectionVersion.version === version;
  });
};

export default async function main({ logger, args, options }) {
  logger.info(`Getting CHANGELOG.md section for ${args.version}`);
  const version = args.version;
  const changelog = await read(resolve(process.cwd(), options.changelog));
  const versionLog = await remark().use(remarkExtractVersionSection, version).process(changelog);
  if (versionLog.messages.length > 0) {
    logger.error(reporter(versionLog.messages));
  }
  // eslint-disable-next-line no-console
  console.log(versionLog.value.replace("\\[", "["));
}

function remarkExtractVersionSection(version) {
  return function transformer(tree, file) {
    const versionSection = find(tree, isVersionSectionForSection(version));
    if (!versionSection) {
      file.fail(`Could not find section for version ${version}`, {
        source: "remark-extract-version-section",
      });
    }
    const previousVersionSection = findAfter(tree, versionSection, isVersionSection);
    const versionChangeSection = previousVersionSection
      ? findAllBetween(tree, versionSection, previousVersionSection)
      : findAllAfter(tree, versionSection);
    return {
      type: "root",
      children: [versionSection, ...versionChangeSection],
    };
  };
}
