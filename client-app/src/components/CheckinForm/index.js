    // @flow
    import React from 'react';
    import { f7 } from 'framework7-react';
    import { List, ListInput, Button, BlockTitle } from 'framework7-react';

    const CheckinForm = () => {
        const saveCheckin = (e: Event) => {
            e.preventDefault();
            var formData = f7.form.convertToData('#checkinform');
        }

        return <List form id="checkinform" onSubmit={saveCheckin}>
            <BlockTitle>Checkin your visit</BlockTitle>
            <ListInput
                label="Name"
                name="name"
                type="text"
                placeholder="Vorname und Nachname"
                clearButton
            />

            <ListInput
                label="Straße und Hausnummer"
                name="street"
                type="text"
                placeholder="Deine Adresse"
                clearButton
            />

            <ListInput
                label="PLZ und Ort"
                name="city"
                type="text"
                placeholder="Deine PLZ inkl. dem Ort"
                clearButton
            />

            <ListInput
                label="Telefonnummer"
                name="phone"
                type="tel"
                placeholder="Deine Telefonnummer"
                clearButton
            />
            <ListInput
                label="E-Mail"
                name="email"
                type="email"
                placeholder="Deine E-Mail-Adresse"
                clearButton
                validate
            />
            <Button type="submit">Submit</Button>
        </List>
    };

    export default CheckinForm;
