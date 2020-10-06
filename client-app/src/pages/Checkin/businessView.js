// @flow
import React from 'react';
import {
    Block,
    BlockTitle,
    Preloader,
    Tab,
    Tabs
} from 'framework7-react';
import i18n from '../../components/i18n.js';
import Logins from '../../components/Logins/index';
import Account from '../../components/Account/index';
import Business from '../../components/Business/index';
import BusinessInfos from '../../components/BusinessInfos/index';
import type {BusinessData, Session} from '../../js/types';
import CheckinForm from '../../components/CheckinForm/index.js';

type Props = {
    session: Session,
    chckrCode: string,
    businessData: BusinessData
}

const BusinessView = (props: Props) => {

    if (props.businessData === null) {
        return <Block className="text-align-center">
            <Preloader color="pink"></Preloader>
        </Block> 
    }

    return <Tabs>
        <Tab id="checkin-ch3ck1n" tabActive>
        <BlockTitle large className="text-align-center">{i18n.t('basic.appname')}</BlockTitle>
                { !props.session.useronline && props.session.connected ?
                    (
                        <Block className="text-align-center">
                            {i18n.t('signin.explanation-short')}
                            <Logins compact={true}></Logins>
                        </Block>
                    ) : null
                }

                { !props.session.connected ? null :
                    (   
                    <Block>
                        <Business businessData={props.businessData}></Business>
                        <CheckinForm></CheckinForm>
                    </Block>
                    )
                }
        </Tab>

        <Tab id="checkin-infos">
        { !props.session.connected ? null :
            (
                <BusinessInfos businessData={props.businessData}></BusinessInfos>
            )
        }
        </Tab>

        <Tab id="checkin-account">
            { !props.session.connected ? null :
                (
                    <Account session={props.session}></Account>
                )
            }
        </Tab>
    </Tabs>
}

export default BusinessView;
