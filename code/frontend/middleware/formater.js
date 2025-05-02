export function formaterContactOfSendInput(contact_input) {
    contact_input.value = contact_input.value.replace(/[^0-9]/g, '')
}


export function formaterContactForShowInput(contact_input) {
    contact_input.value = contactFormaterShow(contact_input.value)
}

export function contactFormaterShow(contact) {
    let formated = contact.replace(/[^0-9]/g, "")

    return formated.replace(/^([0-9]{2})([0-9]{4,5})([0-9]{4})$/, "($1) $2-$3")
}