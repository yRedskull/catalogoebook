import { all_element_link_active, KEY_ALLOWED } from "./constants_global";
import { formaterContactForShowInput } from "./formater";

export function deduplicationListener(object, func, mode = "click") {
    object.removeEventListener(mode, func)
    object.addEventListener(mode, func)
}

export function undisabledForm(form) {
    form.querySelectorAll("input[disabled]").forEach(inp => {
        inp.removeAttribute("disabled")
    });
}

export function capitalizePhrase(phrase) {
    return phrase.split(' ').map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(' ');
}

export function capitalizeFirstLetter(string) {
    return string[0].toUpperCase() + string.slice(1);
}

export function generateObjectElementsInputsForm(form, list_input_name) {
    let inputs = {}

    for (let name of list_input_name) {
        if (form) {
            let selector = form.querySelector(`input[name="${name}"]`)
            if (!selector) {
                selector = form.querySelector(`textarea[name="${name}"]`)
            }
            inputs[name] = selector

        }
    }

    return inputs
}

export function generateObjectElementsOfErrorForm(form, list_input_name, action) {
    let inputs = {}

    for (let name of list_input_name) {
        if (form) {
            inputs[name] = form.querySelector(`#${name}-${action}-error`)
        }
    }

    return inputs
}


export function showConfirmationModal(html_content, cancel_selector_id = '#cancelButton', confirm_selector_id = '#confirmButton') {
    return new Promise((resolve) => {
        const modal = document.createElement('div');
        modal.classList.add('fixed', 'inset-0', 'flex', 'items-center', 'justify-center', 'bg-gray-500', 'bg-opacity-50', 'z-50');
        modal.innerHTML = html_content;
        document.body.appendChild(modal);

        const cancelButton = modal.querySelector(cancel_selector_id);
        const confirmButton = modal.querySelector(confirm_selector_id);

        function handlerCancelBtn() {
            modal.remove();
            resolve(false);
        }
        function handlerConfirmBtn() {
            modal.remove();
            resolve(true);
        }


        cancelButton.removeEventListener('click', handlerCancelBtn);
        cancelButton.addEventListener('click', handlerCancelBtn);

        confirmButton.removeEventListener('click', handlerConfirmBtn);
        confirmButton.addEventListener('click', handlerConfirmBtn);
    });
}

export function preventDefaultLinkActive() {
    for (let element of all_element_link_active) {
        element.removeAttribute("href")

        element.addEventListener("click", e => {
            e.preventDefault()
        })
    }
}

export const contactFormaterWrite = (e) => {
    let target = e.target

    if (writeCondition(e)) return

    formaterContactForShowInput(target)
    if (target.value.replace(/[^0-9]/g, "").length >= 11) return e.preventDefault()
}

export const writeCondition = (e) => e.target.selectionStart != e.target.selectionEnd || KEY_ALLOWED.includes(e.key)

export const nameFormaterWrite = (e) => {
    e.target.value = capitalizePhrase(e.target.value)
}

export const usernameFormaterWrite = (e) => {
    let target = e.target

    if (writeCondition(e)) return
    if (target.value.length >= 50) return e.preventDefault()

    target.value = target.value.toLowerCase()
}


export function addEventListenerForForm(inputs_form, list_input_name, listener_inputs) {
    for (let input_name of list_input_name) {
        for (let [key_listener, func] of Object.entries(listener_inputs[input_name])) {

            try {
                inputs_form[input_name].removeEventListener(key_listener, func)
                inputs_form[input_name].addEventListener(key_listener, func)
            } catch (e) {
                continue
            }

        }
    }
}


export function setNavbarReportColor(nav_link_lead) {
    for (let tag of nav_link_lead) {
        tag.classList.add("bg-green-600")
    }
}

export function clearForm(form) {
    const inputs = form.querySelectorAll("input")
    const textareas = form.querySelectorAll("textarea")

    if (inputs) {
        for (let input of inputs) {
            input.value = ""
        }
    }

    if (textareas) {
        for (let textarea of textareas) {
            textarea.value = ""
        }
    }
}