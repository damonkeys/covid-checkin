// @flow
import React from 'react';
import { Page, Navbar, Block } from 'framework7-react';
import Logo from '../../components/Logo';
import i18n from '../../components/i18n.js';


export default () => (
  <Page>
    <Navbar color="pink" title="chckr"/>
    <Block>
        <Logo direction="horizontal" />
    </Block>
    <Block strong style={{ textAlign: "center" }}>
      <h1>{i18n.t('explanation.bummer')}</h1>
      <h3>{i18n.t('explanation.butcannotusemc')}</h3>
      <h3>{i18n.t('explanation.apologizeemailrequired')}</h3>
    </Block>
    <Block inset>
    <h4>{i18n.t('explanation.mailrecievingprohibited')}</h4>
    <p>{i18n.t('explanation.changemind')}</p>
      <ol>
        <li>{i18n.t('explanation.repair.step1')}</li>
        <li>{i18n.t('explanation.repair.step2')}</li>
        <li>{i18n.t('explanation.repair.step3')}</li>
        <li>{i18n.t('explanation.repair.step4')}</li>
      </ol>
    </Block>
  </Page>
);
