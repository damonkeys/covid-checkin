// @flow
import React from 'react';
import {
    Block,
    BlockTitle,
    List,
    ListItem
} from 'framework7-react';
import Logins from '../../components/Logins/index';
import { useSession } from '../../modules/session';
import { useTranslation } from 'react-i18next';
import type { Session } from '../../js/types';

type Props = {
}

const Account = (props: Props) => {
    const [t] = useTranslation();
    const session: Session = useSession();
    if (session.useronline) {
        return <div>
            <Block>
                <BlockTitle large className="text-align-center">{t('basic.appname')}</BlockTitle>
                <BlockTitle>Profile</BlockTitle>
            </Block>

            <List simple-list>
                <ListItem title="Name" after={session.username}></ListItem>
            </List>
        </div>
        
    }

    return <div>
        <BlockTitle large className="text-align-center">{t('basic.appname')}</BlockTitle>
        <Block>
            {t('signin.explanation')}
        </Block>
        <Logins></Logins>
    </div>
}

export default Account;
