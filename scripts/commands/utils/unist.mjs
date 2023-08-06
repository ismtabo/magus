import { convert } from "unist-util-is";

export const isVersionSection = convert(
  (node) =>
    node.type === "heading" &&
    node.depth === 2 &&
    node.children[0] &&
    /(?:\\)?\[(?<version>.*)\] - \((?<date>.*)\)/.test(node.children[0].value),
);

export function parseVersionHeader(header) {
  const match = header.match(/(?:\\)?\[(?<version>.*)\] - \((?<date>.*)\)/);
  if (!match) {
    return null;
  }
  return {
    version: match.groups.version,
    date: match.groups.date,
  };
}
