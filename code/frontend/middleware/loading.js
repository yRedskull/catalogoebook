export function insertLoading(element) {
    const loading = document.createElement('div')
    {/* <div class="loading">
        <span></span>
        <span></span>
        <span></span>
    </div> */}
    const span_1 = document.createElement('span')
    const span_2 = document.createElement('span')
    const span_3 = document.createElement('span')

    loading.classList.add('loading', 'absolute', "top-1/2", "left-1/2", "transform", "-translate-x-1/2", "-translate-y-1/2", 'z-10')
    loading.appendChild(span_1)
    loading.appendChild(span_2)
    loading.appendChild(span_3)

    element.appendChild(loading)
}

export function removeLoading(element) {
    const loading = element.querySelector('.loading')
    
    if (loading) {
        element.removeChild(loading)
    }
}