export const showModal = (modalContainer, modalContent) => {

  modalContainer.classList.remove("pointer-events-none", "bg-opacity-0");
  modalContainer.classList.add("bg-opacity-50");

  modalContent.classList.remove("opacity-0", "scale-95");
  modalContent.classList.add("opacity-100", "scale-100");
};

export const hideModal = (modalContainer, modalContent) => {
  modalContainer.classList.add("pointer-events-none", "bg-opacity-0");
  modalContainer.classList.remove("bg-opacity-50");

  modalContent.classList.add("opacity-0", "scale-95");
  modalContent.classList.remove("opacity-100", "scale-100");
};



