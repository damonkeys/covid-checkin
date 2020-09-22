// @flow
import React from 'react';
import {
    Block,
    BlockTitle,
} from 'framework7-react';
import i18n from '../../components/i18n.js';
import parse from 'html-react-parser';
import type { BusinessData } from '../../js/types';

type Props = {
    businessData: BusinessData,
}

const BusinessInfos = (props: Props) => {
    if (props.businessData.uuid === undefined) {
        return <Block className="margin-half text-align-center">
            <BlockTitle textColor="red">{i18n.t('business.unknown-code')}</BlockTitle>
        </Block>
    }

    return <div>
        <BlockTitle medium className="no-margin-bottom">{ props.businessData.name }</BlockTitle>
        <BlockTitle grey className="no-margin-top">{ props.businessData.formattedAddress }</BlockTitle>
        { props.businessData.businessInfos && props.businessData.businessInfos.length > 0 ?
            (
                <Block strong className="padding-bottom">{ parse(props.businessData.businessInfos[0].description) }</Block>
            ) : null}
    </div>
}

export default BusinessInfos;
