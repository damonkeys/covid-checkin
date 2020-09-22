// @flow
import React, { useState } from 'react';
import {
    Block,
    Button,
    Input,
    Col, 
    Row
} from 'framework7-react';
import i18n from '../../components/i18n.js';

type Props = {
}

/**
 * Navigation component shows the navigation at the top of a site. It will be visible if a user is logged in only.//#endregion
 * The component connects the backend to give the userdata of a online user. After connecting the backend, the callbackSession prop
 * is called to use usersession-data in other components.
 * 
 * There are some different props:
 * 
 * - callbackSession: it is a callback-function to return the read usersession-data.
 * - hideBacklink: boolean true/false to show or hide the back-link in the navigation bar
 */

const CheckinInput = (props: Props) => {
    const [chckrCode, setChckrCode] = useState('');

    return <Block strong>
        {i18n.t('dashboard.explanation')}<br />
        <h3 className="text-align-center">{i18n.t('dashboard.chckr-code')}</h3>
        <Row>
            <Col width="15"></Col>
            <Col width="70">
                <h1 ><Input inputStyle={{ textAlign: 'center' }} clearButton id="chckr-code" label={i18n.t('dashboard.chckr-code')} type="text" maxlength="5" minlength="5" placeholder={i18n.t('dashboard.chckr-code-placeholder')} onInput={(value) => {setChckrCode(value.currentTarget.value);}}></Input></h1><br />
                <Button raised fill iconF7="checkmark" href={'/checkin/' + chckrCode}>Checkin</Button>
            </Col>
            <Col width="15"></Col>
        </Row>
    </Block>
}

export default CheckinInput;
