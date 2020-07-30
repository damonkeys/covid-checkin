// @flow
import React from 'react';
import { Page, Navbar, Block } from 'framework7-react';
import i18n from '../../components/i18n.js';

export default () => (
  <Page>
    <Navbar color="pink" title="Page not found"/>
    <Block strong>
      <h1>{i18n.t('explanation.bummer')}</h1>
      <h3>{i18n.t('explanation.contentnotfound')}</h3>
    </Block>
  </Page>
);
