// @flow
import React, { useEffect } from 'react';
import { List, ListInput, Button, BlockTitle, f7 } from 'framework7-react';
import checkHTTPError from '../../modules/checkHTTPError';
import { useTranslation } from 'react-i18next';
import type { BusinessProps, UserDataResponse } from '../../js/types';

const UserForm = (props: BusinessProps) => {
    const [t] = useTranslation();
    const { name: businessName, formattedAddress, uuid } = props.businessData;
    const notificationProps = {
        error: {
            icon: '<i class="icon f7-icons color-pink">timer</i>',
            title: t("checkin.userForm.notification.error.title"),
            titleRightText: t("checkin.userForm.notification.error.titleRightText"),
            subtitle: t("checkin.userForm.notification.error.subtitle"),
            text: t("checkin.userForm.notification.error.text"),
            closeButton: true,
            closeOnClick: true
        },
        success: {
            icon: '<i class="icon f7-icons color-pink">hand_thumbsup</i>',
            title: t("checkin.userForm.notification.success.title"),
            titleRightText: t("checkin.userForm.notification.success.titleRightText"),
            subtitle: t("checkin.userForm.notification.success.subtitle"),
            text: t("checkin.userForm.notification.success.text"),
            closeButton: true,
            closeOnClick: true
        }
    };
    const saveCheckin = (e: any) => {
        const formData = f7.form.convertToData('#userForm');
        f7.form.storeFormData('#userForm', formData);
        const { name, street, city, email, phone } = formData;
        fetch('c/checkin', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                user: {
                    name: name,
                    street: street,
                    city: city,
                    country: 'unknown',
                    email: email,
                    phone: phone
                },
                business: {
                    name: businessName,
                    address: formattedAddress,
                    uuid: uuid
                }
            })
        })
            .then(checkHTTPError)
            .then(() => f7.notification.create(notificationProps.success).open())
            .catch((error: Error) => {
                f7.notification.create(notificationProps.error).open();
            });
    };

    useEffect(() => {
        const formData =  f7.form.convertToData('f7form-userForm');// dont ask why this is the name. I don't know.
        f7.form.fillFromData('#userForm',formData);
    });

    useEffect(() => {
        fetch('/c/userdata', {
            method: 'GET'
        })
            .then(response => checkHTTPError(response)) //TODO serverside -> alerting
            .then((response: UserDataResponse) => {
                f7.form.fillFromData('#userForm', {
                    name: response.username,
                    street: response.userstreet,
                    city: response.usercity,
                    country: response.usercountry,
                    email: response.useremail,
                    phone: response.userphone
                });
            })
            .catch((error: number) => {
                //TODO -> show the user the problem (snackbar/toast/whatever)
            });
    }, []);

    // hide checkin user-form if businessData isn't available
    if (!props.businessData.uuid) {
        return <div></div>;
    }

    return <List form id="userForm" formStoreData={true}>
        <BlockTitle>{t("checkin.userForm.title", { bizname: businessName })}</BlockTitle>
        <ListInput
            label={t("checkin.userForm.user.name")}
            name="name"
            type="text"
            placeholder={t("checkin.userForm.user.placeholder")}
            clearButton
        />

        <ListInput
            label={t("checkin.userForm.user.street")}
            name="street"
            type="text"
            placeholder={t("checkin.userForm.user.placeholder")}
            clearButton
        />

        <ListInput
            label={t("checkin.userForm.user.city")}
            name="city"
            type="text"
            placeholder={t("checkin.userForm.user.placeholder")}
            clearButton
        />

        <ListInput
            label={t("checkin.userForm.user.phone")}
            name="phone"
            type="tel"
            placeholder={t("checkin.userForm.user.placeholder")}
            clearButton
        />
        <ListInput
            label={t("checkin.userForm.user.email")}
            name="email"
            type="email"
            placeholder={t("checkin.userForm.user.placeholder")}
            clearButton
            validate
        />
        <Button
            large
            raised
            fill
            iconF7="checkmark"
            onClick={saveCheckin}>{t("checkin.userForm.cta")}</Button>
    </List>
};

export default UserForm;
