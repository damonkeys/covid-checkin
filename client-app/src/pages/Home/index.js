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
import FindBusinessByCheckrCodeForm from '../../components/FindBusinessByCheckrCodeForm/index.js';
import Logo from '../../components/Logo';

const Home = () => {
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
                <FindBusinessByCheckrCodeForm></FindBusinessByCheckrCodeForm>
            </Tab>
        </Tabs>
        <Toolbar tabbar labels bottom>
            <Link tabLink="#home-chckr" iconIos="f7:checkmark_shield" iconAurora="f7:checkmark_shield" iconMd="material:verified_user" tabLinkActive>chckr</Link>
        </Toolbar>
    </Page>
    );
}

export default Home;
