function OpenStufCreator() {
    let updtBtt = document.getElementById('UpdateBtt')
    let crtBtt = document.getElementById('CreateBtt')

    updtBtt.style.display = 'none'
    crtBtt.style.display = 'block'
}

function CreateStuf() {

}

function EditStuf() {

}

function HideShowColor() {
    let colorContainer = document.getElementById('colorContainer')
    let hasColor = document.getElementById('hasColor')

    console.log(hasColor)
    console.log(hasColor.checked)

    if (hasColor.checked) {
        colorContainer.style.display = 'block'
    } else {
        colorContainer.style.display = 'none'
    }

}

function OpenColorCreator() {
    //hide the stuff modal before while going to the colors one
    $('#CreateUpdtModal').modal('hide')

    let colorName = document.getElementById('ColorName')
    let colorRef = document.getElementById('ColorRef')

    colorRef.value = '#000000'
    colorName.value = ''
}

function CloseColorModal() {
    //after we create the color we return to the stuff modal
    $('#CreateColorModal').modal('hide')
    $('#CreateUpdtModal').modal('show')
}

function CreateColor() {
    
    CloseColorModal()
}