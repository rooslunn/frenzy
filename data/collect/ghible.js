// https://www.ghibli.jp/works/red-turtle/

function downloadImage1(src) {
  const link = document.createElement("a")
  link.href = src
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

document.querySelectorAll(".panel-img-top").forEach((img) => {
    downloadImage1(img.src)
})

// 0101010101

async function downloadImage2(imageSrc) {
  const image = await fetch(imageSrc)
  const imageBlog = await image.blob()
  const imageURL = URL.createObjectURL(imageBlog)

  const link = document.createElement("a")
  link.href = imageURL
  link.download = "image file name here"
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

document.querySelectorAll(".panel-img-top").forEach((img) => {
    downloadImage2(img.src)
})