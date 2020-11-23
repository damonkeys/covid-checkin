// @flow
import React from 'react';
import {
    Block,
    BlockTitle,
    Preloader,
    Tab,
    Tabs
} from 'framework7-react';
import Logins from '../../components/Logins/index';
import Account from '../../components/Account/index';
import Business from '../../components/Business/index';
import BusinessInfos from '../../components/BusinessInfos/index';
import type { BusinessProps, Session } from '../../js/types';
import UserForm from '../../components/UserForm/index.js';
import { useSession } from '../../modules/session';
import { useTranslation } from 'react-i18next';


const BusinessView = (props: BusinessProps) => {
    const [t] = useTranslation();
    const session: Session = useSession();

    if (props.businessData === null) {
        return <Block className="text-align-center">
            <Preloader color="pink"></Preloader>
        </Block>
    }


    return <Tabs>
        <Tab id="checkin-ch3ck1n" tabActive>
            <BlockTitle large className="text-align-center">{t('basic.appname')}</BlockTitle>
            {!session.useronline && session.connected ?
                (
                    <Block className="text-align-center">
                        {t('signin.explanation-short')}
                        <Logins compact={true}></Logins>
                    </Block>
                ) : null
            }

            {!session.connected  || !props.businessData.fetched ? null :
                (
                    <Block>
                        <Business businessData={props.businessData}></Business>
                        <UserForm businessData={props.businessData}></UserForm>
                    </Block>
                )
            }
        </Tab>

        <Tab id="checkin-infos">
            {!session.connected || !props.businessData.fetched ? null :
                (
                    <BusinessInfos businessData={props.businessData}></BusinessInfos>
                )
            }
        </Tab>

        <Tab id="checkin-account">
            {!session.connected ? null :
                (
                    <Account></Account>
                )
            }
        </Tab>
    </Tabs>
}

export default BusinessView;