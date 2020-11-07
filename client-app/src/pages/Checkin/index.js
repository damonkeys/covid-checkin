// @flow
import React from 'react';
import {
    Link,
    Navbar,
    NavLeft,
    NavTitle,
    Page,
    Toolbar
} from 'framework7-react';
import BusinessView from '../../components/BusinessView/businessView.js'
import type {BusinessProps } from '../../js/types';
import { useTranslation } from 'react-i18next';



const Checkin = (props: BusinessProps) => {
    const [t] = useTranslation();
    
    return (<Page colorTheme="pink">
        <Navbar color="pink" className="navbar-main">
            <NavLeft>
                <NavLeft backLink={t('basic.back')}></NavLeft>
            </NavLeft>

            <NavTitle>{t('basic.appname')}</NavTitle>
        </Navbar>

        <BusinessView businessData={props.businessData}></BusinessView>

        <Toolbar tabbar labels bottom>
            <Link tabLink="#checkin-ch3ck1n" iconIos="f7:checkmark_shield" iconAurora="f7:checkmark_shield" iconMd="material:verified_user" tabLinkActive>ch3ck1n</Link>
            <Link tabLink="#checkin-infos" iconIos="f7:info_circle" iconAurora="f7:info_circle" iconMd="material:info">Infos</Link>
            <Link tabLink="#checkin-account" iconIos="f7:person_crop_circle" iconAurora="f7:person_crop_circle" iconMd="material:account_circle">Account</Link>
        </Toolbar>
    </Page>);

};

export default Checkin;
