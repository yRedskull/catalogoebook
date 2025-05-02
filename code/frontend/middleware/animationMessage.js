export function showMessage(message, theme) {

    const messageContainer = document.createElement("div");
    messageContainer.className = `
        fixed top-8 left-1/2 transform -translate-x-1/2 z-50
        text-white text-sm font-medium py-4 px-6 rounded-lg
        shadow-lg transition-opacity duration-500 ease-in-out
        opacity-0 ${theme}
    `;
    messageContainer.textContent = message;


    document.body.appendChild(messageContainer);


    setTimeout(() => {
        messageContainer.classList.remove("opacity-0");
        messageContainer.classList.add("opacity-100");
    }, 10);


    setTimeout(() => {
        messageContainer.classList.remove("opacity-100");
        messageContainer.classList.add("opacity-0");

        setTimeout(() => {
            document.body.removeChild(messageContainer);
        }, 500); 
    }, 3000);
}