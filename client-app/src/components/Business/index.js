// @flow
import React from 'react';
import {
    Block,
    BlockTitle
} from 'framework7-react';
import type { BusinessData } from '../../js/types';
import { useTranslation } from 'react-i18next';


type BusinessProps = {
    businessData: BusinessData
}

const Business = (props: BusinessProps) => {
    const [t] = useTranslation();

    if (props.businessData.uuid === undefined) {
        return <Block className="margin-half text-align-center">
            <BlockTitle textColor="red">{t('business.unknown-code', { 'code': props.businessData.code })}</BlockTitle>
        </Block>
    }

    return <Block className="margin-half">
        <BlockTitle medium className="no-margin">{props.businessData.name}</BlockTitle>
        <BlockTitle grey className="no-margin">{props.businessData.formattedAddress}</BlockTitle>
    </Block>
}

export default Business;
