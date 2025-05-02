export function validateName(name) {
    return !/[^a-zA-Z\u00C0-\u017F\s]/u.test(name.trim())
}

export function validateEmail(email) {
    return /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(email)
}

export function validateContact(contact) {
    let value_formated = contact.replace(/[^0-9]/g, '')

    if (value_formated.length != 11) {
        return false
    }

    return /^\d{2}9?\d{8}$/.test(value_formated)
}

export function validateUsername(id) {
    return /^[a-zA-Z0-9]{1,50}$/.test(id)
}

export function validateInput(input_element, validation) {
    if (validation) {
        input_element.classList.add('is-valid')
        input_element.classList.remove('is-invalid')
    } else {
        input_element.classList.add('is-invalid')
        input_element.classList.remove('is-valid')
    }
}

export function verificationFormErrors(inputs_form_for_validate, elements_of_error_form, funcs_validate_form, errors_of_business_rule) {
    try {
        const errors = {}
        let validation

        for (let [_, element_error] of Object.entries(elements_of_error_form)) {
            element_error.textContent = ""
            element_error.classList.add("hidden")
        }

        for (let [validate_name, inp] of Object.entries(inputs_form_for_validate)) {
            if (inp.value == "") {
                continue
            }

            validation = funcs_validate_form[validate_name](inp.value)

            if (!validation) {
                errors[validate_name] = errors_of_business_rule[validate_name]
            }
        }

        if (Object.keys(errors).length > 0) {

            for (let [name_error, value] of Object.entries(errors)) {
                elements_of_error_form[name_error].innerText = value
                elements_of_error_form[name_error].classList.remove("hidden")
            }

            return errors
        }

        return null
    } catch (e) {
        return e
    }
    
}


/* export function validateCPF(cpf) {
    var sum = 0
    var rest

    var strCPF = String(cpf).replace(/[^\d]/g, '')
    
    if (strCPF.length !== 11) return false
    
    if ([
        '00000000000',
        '11111111111',
        '22222222222',
        '33333333333',
        '44444444444',
        '55555555555',
        '66666666666',
        '77777777777',
        '88888888888',
        '99999999999',
        ].indexOf(strCPF) !== -1) return false

    for (let i=1; i<=9; i++){ sum = sum + parseInt(strCPF.substring(i-1, i)) * (11 - i);}

    rest = (sum * 10) % 11

    if ((rest == 10) || (rest == 11)) rest = 0

    if (rest != parseInt(strCPF.substring(9, 10)) ) return false

    sum = 0

    for (let i = 1; i <= 10; i++) {sum = sum + parseInt(strCPF.substring(i-1, i)) * (12 - i)}

    rest = (sum * 10) % 11

    if ((rest == 10) || (rest == 11)) rest = 0

    if (rest != parseInt(strCPF.substring(10, 11) ) ) return false

    return true
} */

/* export function validateDate(date) {
    var data = new Date(date);
    // Verificar se a data é válida
    console.log(data, isNaN(data.getTime()))
    return !isNaN(data.getTime())
      
} */