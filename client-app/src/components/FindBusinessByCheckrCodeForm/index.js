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
 * FindBusinessByCheckrCodeForm is a simple form with a single input field.
 * The user can input a 5-letter "checkr"-Code to find the business she is looking for (to checkin).
 * After touching the button she is redirected to the business page on which she can checkin.
 */
const FindBusinessByCheckrCodeForm = (props: Props) => {
    const [chckrCode, setChckrCode] = useState('');

    return <Block strong>
        {i18n.t('dashboard.explanation')}<br />
        <h3 className="text-align-center">{i18n.t('dashboard.chckr-code')}</h3>
        <Row>
            <Col width="15"></Col>
            <Col width="70">
                <h1 >
                    <Input
                    inputStyle={{ textAlign: 'center' }}
                    clearButton id="chckr-code"
                    label={i18n.t('dashboard.chckr-code')}
                    type="text"
                    maxlength="5"
                    minlength="2"
                    placeholder={i18n.t('dashboard.chckr-code-placeholder')}
                    onInput={(value) => {setChckrCode(value.currentTarget.value);}}>
                    </Input>
                </h1>
                <br />
                <Button large raised fill iconF7="checkmark" href={'/checkin/' + chckrCode}>{i18n.t('business.findBizAndCheckin')}</Button>
            </Col>
            <Col width="15"></Col>
        </Row>
    </Block>
}

export default FindBusinessByCheckrCodeForm;
