// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');
const { default: dedent } = require('ts-dedent');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Magus',
  tagline: 'MAgic Generator Utility for Software',
  favicon: 'img/favicon.ico',

  // Set the production url of your site here
  url: 'https://ismtabo.github.io/',
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: '/magus/',

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'ismtabo', // Usually your GitHub org/user name.
  projectName: 'magus', // Usually your repo name.

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/ismtabo/magus/tree/main/docs/page/',
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/ismtabo/magus/tree/main/docs/page/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      // TODO: Replace with your project's social card
      // image: 'img/docusaurus-social-card.jpg',
      metadata: [
        { name: 'keywords', content: 'magus, go, golang, cli, code generation, code generator, codegen, codegen tool, codegen utility, codegen cli, codegen command line tool, codegen command line utility, codegen cli tool, codegen cli utility, codegen command line' },
      ],
      navbar: {
        title: 'Magus',
        logo: {
          alt: 'Magus Logo',
          src: 'img/sage.png',
        },
        items: [
          {
            type: 'docSidebar',
            sidebarId: 'docsSidebar',
            position: 'left',
            label: 'Docs',
          },
          // {to: '/blog', label: 'Blog', position: 'left'},
          {
            type: 'docsVersionDropdown',
            position: 'right',
            dropdownItemsAfter: [{to: '/versions', label: 'All versions'}],
            dropdownActiveClassDisabled: true,
          },
          {
            href: 'https://pkg.go.dev/github.com/ismtabo/magus/v2',
            label: 'GoDoc',
            position: 'right',
          },
          {
            href: 'https://github.com/ismtabo/magus',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Docs',
            items: [
              {
                label: 'Tutorial',
                to: '/docs/intro',
              },
            ],
          },
          {
            title: 'More',
            items: [
              {
                label: 'GoDoc',
                href: 'https://pkg.go.dev/github.com/ismtabo/magus/v2'
              },
              {
                label: 'GitHub',
                href: 'https://github.com/facebook/docusaurus',
              },
            ],
          },
        ],
        copyright: dedent`
        Copyright © ${new Date().getFullYear()} Ismael Taboada. Built with Docusaurus.<br>
        Awesome gopher images from <a href="https://github.com/ashleymcnamara/gophers">ashleymcnamara/gophers</a>.
        `,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
  customFields: {
    description: 'Magus is a CLI tool that generates code for your projects.',
    releasesPage: `https://github.com/ismtabo/magus/releases`
  }
};

module.exports = config;
