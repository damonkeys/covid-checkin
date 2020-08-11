// @flow
import React, { useEffect } from 'react';
import ActivationInformation from '../../components/Activationinformation';

type Props = {
    code: string,
    activationToken: string
}

const Activation = (props: Props) => {

    const basicTokenValidation = (activationToken: string): boolean => {
        if (activationToken) {
            if (activationToken.length === 36) {
                if (activationToken.split('-').length === 5) {
                    return true;
                }
            }
        }
        return false;
    }

    const fetchDataOnValidToken = () => {
        if (basicTokenValidation(props.activationToken)) {
            fetch(`/auth/activation/${props.activationToken}`, {
                method: 'GET'
            })
        }
    }

    useEffect(fetchDataOnValidToken);


    const renderContent = () => {
        if (props.activationToken === 'success') {
            // return (
            //     <Page>
            //         <BlockTitle>{i18n.t('activation.success.title')}</BlockTitle>
            //         <Block strong>
            //             <Link href="/" external>{i18n.t('activation.toStartPage')} </Link>
            //         </Block>
            //     </Page>)
            return <ActivationInformation activationState="success"></ActivationInformation>
        } else {
            return <ActivationInformation activationState="ongoing"></ActivationInformation>
        }
    }

    return renderContent();
}

export default Activation;
