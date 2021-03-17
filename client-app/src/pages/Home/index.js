// @flow
import React from 'react';
import {
    Block,
    Link,
    Navbar,
    NavLeft,
    NavTitle,
    Page,
    Tabs,
    Tab,
    Toolbar
} from 'framework7-react';
import { f7 } from 'framework7-react';
import i18n from '../../components/i18n.js';
import Logins from '../../components/Logins/index';
import FindBusinessByCheckrCodeForm from '../../components/FindBusinessByCheckrCodeForm/index.js';
import Account from '../../components/Account';
import Logo from '../../components/Logo';
import { useSession } from '../../modules/session';
import type { Session } from '../../js/types';

const Home = () => {
    const session: Session = useSession();
    return (<Page colorTheme="pink">
        <Navbar color="pink" className="navbar-main">
            <NavLeft>
                {(f7.views.main.router.history.length <= 1) ? null :
                    (
                        <NavLeft backLink={i18n.t('basic.back')}></NavLeft>
                    )
                }
            </NavLeft>

            <NavTitle>{i18n.t('basic.appname')}</NavTitle>
        </Navbar>
        <Tabs>
            <Tab id="home-chckr" tabActive>
                <Block>
                    <Logo direction="horizontal" />
                </Block>
  
                { session.useronline || !session.connected ? null :
                    (
                        <Block className="text-align-center">
                            {i18n.t('signin.explanation-short')}
                            <Logins compact={true}></Logins>
                        </Block>
                    )
                }
                <FindBusinessByCheckrCodeForm></FindBusinessByCheckrCodeForm>
            </Tab>

            <Tab id="home-account">
                <Account session={session}></Account>
            </Tab>
        </Tabs>
        <Toolbar tabbar labels bottom>
            <Link tabLink="#home-chckr" iconIos="f7:checkmark_shield" iconAurora="f7:checkmark_shield" iconMd="material:verified_user" tabLinkActive>chckr</Link>
            <Link tabLink="#home-account" iconIos="f7:person_crop_circle" iconAurora="f7:person_crop_circle" iconMd="material:account_circle">Account</Link>
        </Toolbar>
    </Page>
    );
}

export default Home;
