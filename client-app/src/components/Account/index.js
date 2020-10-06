// @flow
import React from 'react';
import {
    Block,
    BlockTitle,
    List,
    ListItem
} from 'framework7-react';
import Logins from '../../components/Logins/index';
import i18n from '../../components/i18n.js';
import type { Session } from '../../js/types';

type Props = {
    session: Session,
}

const Account = (props: Props) => {

    if (props.session.useronline) {
        return (
            <div>
                <Block>
                    <BlockTitle large className="text-align-center">{i18n.t('basic.appname')}</BlockTitle>
                    <BlockTitle>Profile</BlockTitle>
                </Block>

                <List simple-list>
                    <ListItem title="Name" after={props.session.username}></ListItem>
                </List>
            </div>
        )
    }

    return (
        <div>
            <BlockTitle large className="text-align-center">{i18n.t('basic.appname')}</BlockTitle>
            <Block>
                {i18n.t('signin.explanation')}
            </Block>
            <Logins></Logins>
        </div>
    )
}

export default Account;
