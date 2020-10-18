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
import type { BusinessData } from '../../js/types';
import CheckinForm from '../../components/CheckinForm/index.js';
import { getSession } from '../../modules/session';
import { useTranslation } from 'react-i18next';

type BusinessViewProps = {
    businessData: BusinessData
}

const BusinessView = (props: BusinessViewProps) => {
    const [t] = useTranslation();

    if (props.businessData === null) {
        return <Block className="text-align-center">
            <Preloader color="pink"></Preloader>
        </Block>
    }

    const session = getSession();

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

            {!session.connected ? null :
                (
                    <Block>
                        <Business businessData={props.businessData}></Business>
                        <CheckinForm></CheckinForm>
                    </Block>
                )
            }
        </Tab>

        <Tab id="checkin-infos">
            {!session.connected ? null :
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
