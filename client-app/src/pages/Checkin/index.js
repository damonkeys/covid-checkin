// @flow
import React, { useState, useEffect } from 'react';
import { f7 } from 'framework7-react';
import {
    Link,
    Navbar,
    NavLeft,
    NavTitle,
    Page,
    Toolbar
} from 'framework7-react';
import i18n from '../../components/i18n.js';
import BusinessView from './businessView.js'
import type { BusinessData } from '../../js/types';
import { checkHTTPError } from '../../modules/error';
import { getSession, updateSession } from '../../modules/session';

type Props = {
    chckrCode: string
}

const Checkin = (props: Props) => {

    const [businessData: BusinessData, setBusinessData] = useState({});

    useEffect(() => {
        updateSession();
    }, []);

    useEffect(() => {
        if (props.chckrCode === undefined) {
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
                setBusinessData({
                    code: props.chckrCode
                });
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
        var address: string[] = [];
        address.push(businessData.street || '');
        address.push(businessData.city || '');
        address = address.filter(function (el) {
            return el !== undefined && el.trim() !== '';
        });
        return address.join(', ');
    }


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

        {(f7.views.main.router.currentRoute && f7.views.main.router.currentRoute.params) ?
            (
                <BusinessView chckrCode={f7.views.main.router.currentRoute.params.chckrCode} session={getSession()} businessData={businessData}></BusinessView>
            ) : null
        }

        <Toolbar tabbar labels bottom>
            <Link tabLink="#checkin-ch3ck1n" iconIos="f7:checkmark_shield" iconAurora="f7:checkmark_shield" iconMd="material:verified_user" tabLinkActive>ch3ck1n</Link>
            <Link tabLink="#checkin-infos" iconIos="f7:info_circle" iconAurora="f7:info_circle" iconMd="material:info">Infos</Link>
            <Link tabLink="#checkin-account" iconIos="f7:person_crop_circle" iconAurora="f7:person_crop_circle" iconMd="material:account_circle">Account</Link>
        </Toolbar>
    </Page>);

}

export default Checkin;
