// @flow
import React, { useState, useEffect } from 'react';
import {
    Block,
    Button,
    Input,
    Col,
    Row
} from 'framework7-react';
import type { BusinessData } from '../../js/types';
import checkHTTPError from '../../modules/checkHTTPError';
import { useTranslation } from 'react-i18next';

type FindBusinessByCheckrCodeFormProps = {
}

/**
 * FindBusinessByCheckrCodeForm is a simple form with a single input field.
 * The user can input a 5-letter "checkr"-Code to find the business she is looking for (to checkin).
 * After touching the button she is redirected to the business page on which she can checkin.
 */
const FindBusinessByCheckrCodeForm = (props: FindBusinessByCheckrCodeFormProps) => {
    const [t] = useTranslation();
    const initialInfoText = t("checkin.initialInfo");
    const [chckrCode: string, setChckrCode] = useState('');
    const [businessData: BusinessData, setBusinessData] = useState({});
    const [infoMessage: string, setInfoMessage] = useState(initialInfoText);

    useEffect(() => {
        if (chckrCode.length < 5) {
            return;
        }
        fetch('/biz/business/' + chckrCode, {
            method: 'GET'
        })
            .then(checkHTTPError)
            .then((response: BusinessData) => {
                setInfoMessage(t("checkin.validCode"));
                response.formattedAddress = formatAddress(response);
                response.fetched = true;
                setBusinessData(response);
            })
            .catch((error: number) => {
                setInfoMessage(t("checkin.invalidCode"));
                setBusinessData({});
            });
    }, [chckrCode, t]);

    // this function extracts out of BusinessData-response all address-data and returns an
    // address-string for showing.
    //
    // It returns this format:
    // <street>, <city>
    //
    // if there es one element undefined it won't be in the result string.
    const formatAddress = (businessData: BusinessData): string => {
        var address: string[] = [];
        address.push(businessData.street || '');
        address.push(businessData.city || '');
        address = address.filter(function (el) {
            return el !== undefined && el.trim() !== '';
        });
        return address.join(', ');
    };

    const checkinPossible = (): boolean => {
        return (chckrCode.length < 5 || businessData.uuid === undefined);
    };

    const runInputCleared = (event: any) => {
        setInfoMessage(initialInfoText);
    };

    return <Block strong>
        {t('dashboard.explanation')}<br />
        <h3 className="text-align-center">{t('dashboard.chckr-code')}</h3>
        <Row>
            <Col width="15"></Col>
            <Col width="70">
                <h1 style={{ textAlign: 'center' }}>
                    <Input
                        inputStyle={{ textAlign: 'center' }}
                        clearButton id="chckr-code"
                        label={t('dashboard.chckr-code')}
                        type="text"
                        maxlength="5"
                        placeholder={t('dashboard.chckr-code-placeholder')}
                        onInput={(value) => { setChckrCode(value.currentTarget.value); }}
                        validate={false}
                        info={infoMessage}
                        onInputClear={runInputCleared}
                        onInputEmpty={runInputCleared}
                    >
                    </Input>
                </h1>
                <br />
                <Button
                    large
                    raised
                    fill
                    iconF7="checkmark"
                    href={'/checkin/' + chckrCode}
                    preventRouter={checkinPossible()}
                    routeProps={{ businessData: businessData }}>{t('business.findBizAndCheckin')}
                </Button>
            </Col>
            <Col width="15"></Col>
        </Row>
    </Block>
}

export default FindBusinessByCheckrCodeForm;
