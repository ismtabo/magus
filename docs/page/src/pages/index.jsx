import React from 'react';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import HomepageFeatures from '@site/src/components/HomepageFeatures';
import logo from '@site/static/img/DOCTOR_STRANGE_GOPHER.png';
import Codeblock from '@theme/CodeBlock';
import Layout from '@theme/Layout';
import clsx from 'clsx';

import styles from './index.module.css';

function HomepageHeader() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <h1 className="hero__title">{siteConfig.title}</h1>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <p><img className={styles.heroImage} src={logo} alt={siteConfig.title} /></p>
        <div className="container">
          <div className={clsx("row row--align-center", styles[`row--justify-center`])}>
            <div className="col col--5">
              <Codeblock language="sh" >
                {`go install github.com/ismtabo/magus@latest`}
              </Codeblock>
            </div>
          </div>
        </div>
        <div className={styles.buttons}>
          <Link
            className="button button--secondary button--lg"
            to="/docs/intro">
            Magus Tutorial - 5min ⏱️
          </Link>
          <Link
            className="button button--info button--lg"
            href={`${siteConfig.customFields.releasesPage}`}>
            Releases
          </Link>
        </div>
      </div>
    </header>
  );
}

export default function Home() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title={`Hello from ${siteConfig.title}`}
      description={siteConfig.customFields.description}>
      <HomepageHeader />
      <main>
        <HomepageFeatures />
      </main>
    </Layout>
  );
}
