import React from 'react';
import clsx from 'clsx';
import styles from './styles.module.css';

/** 
 * @typedef {object} Feature 
 * @property {string} title
 * @property {string} description
 */

/** @type {Feature[]} */
const FeatureList = [
  {
    title: 'Easy to Use',
    description: (
      <>
        Magus is designed from the ground up to be easily installed and used to
        generate source code for your projects.
      </>
    ),
  },
  {
    title: 'Zero Configuration',
    description: (
      <>
        Magus requires no configuration. Just install it,
        write your own templates and start generating code.
      </>
    ),
  },
];

/** @param {Feature} options */
function Feature({ title, description }) {
  return (
    <div className={clsx('col col--6')}>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
