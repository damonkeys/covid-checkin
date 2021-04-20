// @flow
import React from 'react';
import {
    Block,
    BlockTitle,
    List,
    ListItem
} from 'framework7-react';
import Register from '../../components/Register/index';
import Logo from '../../components/Logo';
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
                <Logo direction="horizontal" />
            </Block>
            <Block>
                <BlockTitle>Profile</BlockTitle>
            </Block>

            <List simple-list>
                <ListItem title="Name" after={session.username}></ListItem>
            </List>
        </div>
        
    }

    return <div>
        <Block>
            <Logo direction="horizontal" />
        </Block>
        <Block>
            {t('signin.explanation')}
        </Block>
        <Register></Register>
    </div>
}

export default Account;
