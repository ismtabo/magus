import React from 'react';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';

/** 
 * @param {object} options
 * @param {string} options.title
 */
export default function ReleasesPageLink({ title }) {
  const { siteConfig } = useDocusaurusContext();
  return (
    <a href={`${siteConfig.customFields.releasesPage}`}>{title}</a>
  );
}