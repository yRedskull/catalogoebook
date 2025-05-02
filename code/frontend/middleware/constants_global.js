export const ID_STRING = "id"
export const NAME_STRING = "name"
export const CONTACT_STRING = "contact"
export const ID_ATT_STRING = `${ID_STRING}-attendant`
export const NAME_ATT_STRING = `${NAME_STRING}-attendant`
export const CONTACT_ATT_STRING = `${CONTACT_STRING}-attendant`
export const MAX_LEAD_ATT_STRING = "max-lead-attendant"
export const USERNAME_STRING = "username"
export const EMAIL_STRING = "email"
export const PASSWORD_STRING = "password"
export const CONFIRM_PASSWORD_STRING = "confirm-password"
export const CREATE_STRING = "create"
export const EDIT_STRING = "edit"
export const ERROR_STRING = "error"
export const ID_MESS_STRING = `${ID_STRING}-message`
export const LEAD_STRING = "lead"
export const MESSAGE_STRING = "message"

export const REGEX_WTS = /^([0-9]{2})([0-9]{4,5})([0-9]{4})$/
export const KEY_ALLOWED = ['ArrowLeft', 'Backspace', 'ArrowRight', 'Control', 'Shift']
export const MAX_LEAD_LIMIT = 10000
export const list_input_name_attendant = [ID_ATT_STRING, NAME_ATT_STRING, CONTACT_ATT_STRING, MAX_LEAD_ATT_STRING]
export const list_input_name_message = [ID_MESS_STRING, MESSAGE_STRING]
export const list_input_name_lead = [NAME_STRING, CONTACT_STRING]


export const errors_of_business_rule = new Object()
errors_of_business_rule[ID_STRING] = errors_of_business_rule[ID_ATT_STRING] = "O ID não deve conter símbolos ou espaços"
errors_of_business_rule[NAME_STRING] = errors_of_business_rule[NAME_ATT_STRING] = "O nome não deve conter símbolos ou números"
errors_of_business_rule[CONTACT_STRING] = errors_of_business_rule[CONTACT_ATT_STRING] = "O contato deverá ser nesse formato (DDD) 9XXXX-XXXX"
errors_of_business_rule[MAX_LEAD_ATT_STRING] = "O número de redirecionamentos deverá ser entre 1-10000"
errors_of_business_rule[USERNAME_STRING] = "Deverá conter apenas letras ou números sem espaços"

errors_of_business_rule[EMAIL_STRING] = "Email inválido"
errors_of_business_rule[CONFIRM_PASSWORD_STRING] = "Senhas diferentes"

export const full_url = new URL(window.location.href)
export const pathname_host = full_url.pathname;
export const search_url = full_url.search
export const url_without_protocol = pathname_host + search_url
export const all_element_link_active = document.querySelectorAll(`a[href="${pathname_host}"]`)
