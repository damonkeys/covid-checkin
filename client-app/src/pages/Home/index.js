// @flow
import React from 'react';
import {
    Block,
    BlockTitle,
    Link,
    Page,
    Tabs,
    Tab,
    Toolbar
} from 'framework7-react';
import i18n from '../../components/i18n.js';
import SessionComponent from '../../components/SessionComponent/index';
import Logins from '../../components/Logins/index';
import Navigation from '../../components/Navigation/index';
import CheckinInput from '../../components/CheckinInput/index.js';
import Account from '../../components/Account';

export default class Home extends SessionComponent {
    render() {
        return (
            <Page colorTheme="pink">
                <Navigation />
                <Toolbar tabbar labels bottom>
                    <Link tabLink="#home-ch3ck1n" iconIos="f7:checkmark_shield" iconAurora="f7:checkmark_shield" iconMd="material:verified_user" tabLinkActive>ch3ck1n</Link>
                    <Link tabLink="#home-account" iconIos="f7:person_crop_circle" iconAurora="f7:person_crop_circle" iconMd="material:account_circle">Account</Link>    
                </Toolbar>
                
                <Tabs swipeable>
                    <Tab id="home-ch3ck1n" tabActive>
                        <BlockTitle large className="text-align-center">{i18n.t('basic.appname')}</BlockTitle>
                        { (this.state.session.useronline || !this.state.session.connected) ? (null) :
                            (
                                <Block className="text-align-center">
                                    {i18n.t('signin.explanation-short')}
                                    <Logins compact={true}></Logins>
                                </Block>
                            )
                        }

                        { !this.state.session.connected ? (null) :
                            (
                                <CheckinInput></CheckinInput>
                            )
                        }
                    </Tab>

                    <Tab id="home-account">
                        { !this.state.session.connected ? (null) :
                            (
                                <Account session={this.state.session}></Account>
                            )
                        }
                    </Tab>
                </Tabs>
            </Page>            
        );
    }
}
