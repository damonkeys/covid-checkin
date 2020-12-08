    // @flow
    import React, { useState } from 'react';
    import { Page, Navbar, Block, Button, Popup, NavRight, Link, Card, CardHeader, CardContent, List, ListInput } from 'framework7-react';
    import Logo from '../../components/Logo';
    import i18n from '../i18n';

    type Props = {
        activationState: 'ongoing' | 'success'
    }
    const ActivationInformation = (props: Props, { f7router }:any) => {

        const [isPopUpVisible, setPopUpVisible] = useState(false);
        const [email, setEmail] = useState('');

        const fetchI18NKey = (translationKey: string) => {
            return `activation.${props.activationState}.${translationKey}`;
        }

        const popUpOrCheckin = ():React.Element => {
            if(props.activationState === 'success') {
                return  <Button large fill raised href='/'>{i18n.t(i18n.t(fetchI18NKey('action')))}</Button>
            }
            return <Button large fill raised onClick={(event) => setPopUpVisible(isPopUpVisible)}>{i18n.t(i18n.t(fetchI18NKey('action')))}</Button>

        }

        const emailValid = (): boolean => {
            // poor mans validation. Smarter (not better) validation would improve UX. -> help the user on typos
            if (email) {
                if (email.indexOf('@') !== -1) {
                    if (email.indexOf('.') !== -1) {
                        return true;
                    }
                }
            }
            return false;
        }

        const resendActivationMail = () => {
            fetch(`/auth/activation/${email}`, {
                method: 'POST'
            })
                .then((response) => response.json())
                .then(async (json) => {
                    console.log('###### :' + json);
                });
        }

        const sendMailOnValidEmail = () => {
            if (emailValid()) {
                resendActivationMail();
                // todo - show positive outcome
            }
        }

        return <Page colorTheme="pink">
            <Navbar color="pink" title={i18n.t(fetchI18NKey('head'))} />
            <Block>
                <Logo direction="horizontal" />
            </Block>
            <Block strong>
                <h2>{i18n.t(fetchI18NKey('title'))}</h2>
                <h3>{i18n.t(fetchI18NKey('text'))}</h3>
                <Block>
                    {popUpOrCheckin()}
                </Block>
            </Block>
            <Popup colorTheme="pink" opened={isPopUpVisible} onPopupClosed={() => setPopUpVisible(!isPopUpVisible)}>
                <Page>
                    <Navbar title={i18n.t('activation.resendActivation.title')}>
                        <NavRight>
                            <Link popupClose>{i18n.t('basic.close')}</Link>
                        </NavRight>
                    </Navbar>
                    <Card className="card">
                        <CardHeader className="no-border">
                            <div>{i18n.t('activation.resendActivation.intro')}</div>
                        </CardHeader>
                        <CardContent padding={false}>
                            <List noHairlinesMd>
                                <ListInput
                                    label="E-mail"
                                    floatingLabel
                                    type="email"
                                    placeholder="Your e-mail"
                                    clearButton
                                    required
                                    validate
                                    onChange={(event) => setEmail(event.target.value)}
                                ></ListInput></List>
                        </CardContent>
                    </Card>
                    <Block>
                        <Button large fill raised onClick={(event) => sendMailOnValidEmail()}>{i18n.t(i18n.t(fetchI18NKey('action')))}</Button>
                    </Block>
                </Page>
            </Popup>
        </Page>


    };

    export default ActivationInformation;
