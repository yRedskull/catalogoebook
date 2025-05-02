import { preventDefaultLinkActive } from "./utils_global";

!function() {
    const hamburgerButton = document.getElementById('hamburger');
    const mobileMenu = document.getElementById('mobile-menu');

    hamburgerButton.addEventListener('click', function () {
        this.classList.toggle('open');

        if (mobileMenu) {
            mobileMenu.classList.toggle('hidden');
            mobileMenu.classList.toggle('flex')
        }
    });
    
    preventDefaultLinkActive()
}()