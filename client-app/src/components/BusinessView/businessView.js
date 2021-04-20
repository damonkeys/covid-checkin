// @flow
import React from 'react';
import {
    Block,
    Preloader,
    Tab,
    Tabs
} from 'framework7-react';
import Logo from '../../components/Logo';
import Business from '../../components/Business/index';
import BusinessInfos from '../../components/BusinessInfos/index';
import type { BusinessProps } from '../../js/types';
import UserForm from '../../components/UserForm/index.js';

const BusinessView = (props: BusinessProps) => {
    if (props.businessData === null) {
        return <Block className="text-align-center">
            <Preloader color="pink"></Preloader>
        </Block>
    }


    return <Tabs>
        <Tab id="checkin-chckr" tabActive>
            <Block>
                <Logo direction="horizontal" />
            </Block>
            {!props.businessData.fetched ? null :
                (
                    <Block>
                        <Business businessData={props.businessData}></Business>
                        <UserForm businessData={props.businessData}></UserForm>
                    </Block>
                )
            }
        </Tab>

        <Tab id="checkin-infos">
            {!props.businessData.fetched ? null :
                (
                    <BusinessInfos businessData={props.businessData}></BusinessInfos>
                )
            }
        </Tab>
    </Tabs>
}

export default BusinessView;
