import React from 'react';
import clsx from 'clsx';

import styles from './styles.module.css';

/** 
 * @typedef {object} Feature 
 * @property {string} title
 * @property {string} description
 * @property {string} image
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
    image: require('@site/static/img/GOPHER_LAPTOP.png').default,
  },
  {
    title: 'Zero Configuration',
    description: (
      <>
        Magus requires no additional configuration. Just install it,
        write your own templates and start generating code.
      </>
    ),
    image: require('@site/static/img/GOPHER_MIC_DROP.png').default,
  },
];

/** @param {Feature} options */
function Feature({ title, description, image }) {
  return (
    <div className={clsx('col col--6')}>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
        <p><img className={styles.featureImg} src={image} alt={title} /></p>
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
