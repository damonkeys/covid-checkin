// @flow
import React, {useState, useEffect} from 'react';
import {
  Link,
  Navbar,
  NavLeft,
  NavTitle,
  Page,
  Toolbar
} from 'framework7-react';
import checkHTTPError from '../../modules/checkHTTPError';
import BusinessView from '../../components/BusinessView/businessView.js'
import { useSession } from '../../modules/session';
import type { CheckinProps, BusinessData } from '../../js/types';
import {useTranslation} from 'react-i18next';

const Checkin = (props: CheckinProps) => {
  const [t] = useTranslation();
  const [businessData: BusinessData, setBusinessData] = useState({});
  useSession();


  useEffect(() => {
    if (props.businessData) {
      setBusinessData(props.businessData);
    }
    // no data given by props, then do a server-fetch
    fetch('/biz/business/' + props.chckrCode, {method: 'GET'})
      .then(checkHTTPError)
      .then((response : BusinessData) => {
        response.formattedAddress = formatAddress( response );
        response.fetched = true;
        setBusinessData(response);
      })
      .catch((error : number) => {
        setBusinessData({fetched: true});
      }
    );
  }, [props]);

  
  // this function extracts out of BusinessData-response all address-data and
  // returns an address-string for showing.
  //
  // It returns this format: <street>, <city>
  //
  // if there es one element undefined it won't be in the result string.
  const formatAddress = (businessData : BusinessData) : string => {
    var address : string[] = [];
    address.push(businessData.street || '');
    address.push(businessData.city || '');
    address = address.filter(function (el) {
      return el !== undefined && el.trim() !== '';
    });
    return address.join(', ');
  };

  return (
    <Page colorTheme="pink">
      <Navbar color="pink" className="navbar-main">
        <NavLeft>
          <NavLeft backLink={t('basic.back')}></NavLeft>
        </NavLeft>

        <NavTitle>{t('basic.appname')}</NavTitle>
      </Navbar>

      <BusinessView businessData={ businessData }></BusinessView>

      <Toolbar tabbar labels bottom>
        <Link
          tabLink="#checkin-ch3ck1n"
          iconIos="f7:checkmark_shield"
          iconAurora="f7:checkmark_shield"
          iconMd="material:verified_user"
          tabLinkActive>ch3ck1n</Link>
        <Link
          tabLink="#checkin-infos"
          iconIos="f7:info_circle"
          iconAurora="f7:info_circle"
          iconMd="material:info">Infos</Link>
        <Link
          tabLink="#checkin-account"
          iconIos="f7:person_crop_circle"
          iconAurora="f7:person_crop_circle"
          iconMd="material:account_circle">Account</Link>
      </Toolbar>
    </Page>
  );

};

export default Checkin;
