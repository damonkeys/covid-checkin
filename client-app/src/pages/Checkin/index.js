// @flow
import React from 'react';
import { f7 } from 'framework7-react';
import {
    Link,
    Navbar,
    NavLeft,
    NavTitle,
    Page,
    Toolbar
} from 'framework7-react';
import SessionComponent from '../../components/SessionComponent/index';
import i18n from '../../components/i18n.js';
import TabContent from './tabContent.js'

export default class Checkin extends SessionComponent {
    render() {
        return (
            <Page colorTheme="pink">
                <Navbar color="pink" className="navbar-main">
                    <NavLeft>
                        {( f7.views.main.router.history.length <= 1) ? null :
                            (
                                <NavLeft backLink={i18n.t('basic.back')}></NavLeft>
                            )
                        }
                    </NavLeft>

                    <NavTitle>{i18n.t('basic.appname')}</NavTitle>
                </Navbar>
                
                {( f7.views.main.router.currentRoute && f7.views.main.router.currentRoute.params) ? 
                    (
                        <TabContent chckrCode={f7.views.main.router.currentRoute.params.chckrCode} session={this.state}></TabContent>
                    ) : null
                }

                <Toolbar tabbar labels bottom>
                    <Link tabLink="#checkin-ch3ck1n" iconIos="f7:checkmark_shield" iconAurora="f7:checkmark_shield" iconMd="material:verified_user" tabLinkActive>ch3ck1n</Link>
                    <Link tabLink="#checkin-infos" iconIos="f7:info_circle" iconAurora="f7:info_circle" iconMd="material:info">Infos</Link>
                    <Link tabLink="#checkin-account" iconIos="f7:person_crop_circle" iconAurora="f7:person_crop_circle" iconMd="material:account_circle">Account</Link>
                </Toolbar>
            </Page>
        );
    }
}
