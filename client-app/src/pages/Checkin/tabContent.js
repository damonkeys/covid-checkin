// @flow
import React, {useState, useEffect} from 'react';
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
import { checkHTTPError } from '../../modules/error';
import type {BusinessData, Session} from '../../js/types';

type Props = {
    session: Session,
    chckrCode: string,
}

const TabContent = (props: Props) => {
    const [businessData: BusinessData, setBusinessData] = useState(null);

    useEffect(() => {
        if(props.chckrCode === undefined) {
            return;
        }
        fetch('/biz/business/' + props.chckrCode, {
            method: 'GET'
        })
            .then(checkHTTPError)
            .then((response: BusinessData) => {
                response.formattedAddress = formatAddress(response);
                setBusinessData(response);
            })
            .catch((error: number) => {
                setBusinessData({});
            });
    }, [props.chckrCode]);

    // this function extracts out of BusinessData-response all address-data and returns an
    // address-string for showing.
    //
    // It returns this format:
    // <street>, <city>
    //
    // if there es one element undefined it won't be in the result string.
    const formatAddress = (businessData: BusinessData) => {
        var address:string[] = [];
        address.push(businessData.street || '');
        address.push(businessData.city || '');
        address = address.filter(function (el) {
            return el !== undefined && el.trim() !== '';
        });
        return address.join(', ');
    }

    if (businessData === null) {
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
                        <Business businessData={businessData}></Business>
                    )
                }
        </Tab>

        <Tab id="checkin-infos">
        { !props.session.connected ? null :
            (
                <BusinessInfos businessData={businessData}></BusinessInfos>
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

export default TabContent;
