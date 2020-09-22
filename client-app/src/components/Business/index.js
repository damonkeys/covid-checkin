// @flow
import React from 'react';
import {
    Block,
    BlockTitle
} from 'framework7-react';
import CheckinForm from '../CheckinForm/index.js';
import i18n from '../../components/i18n.js';
import type { BusinessData } from '../../js/types';


type Props = {
    businessData: BusinessData,
}

const Business = (props: Props) => {
    if (props.businessData.uuid === undefined) {
        return <Block className="margin-half text-align-center">
                <BlockTitle textColor="red">{i18n.t('business.unknown-code')}</BlockTitle>
            </Block>
    }

    return <div>
                <Block className="margin-half">
                    <BlockTitle medium className="no-margin">{props.businessData.name}</BlockTitle>
                    <BlockTitle grey className="no-margin">{props.businessData.formattedAddress}</BlockTitle>
                </Block>
                <Block strong>
                    <CheckinForm></CheckinForm>
                </Block>
            </div>
}

export default Business;
