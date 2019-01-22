export function isValidFileType(file, ...filetype) {
  if (!file instanceof FileList) return false;
  if (file[0] == null) return false;
  const regex = new RegExp(`[\\s\\S]+[.]+${generateRegex(filetype)}$`)
  return regex.test(file[0].name)
}

function generateRegex(fileType) {
  let str = "";
  fileType.forEach(value => {
    str = `${str}|${value}`
  })
  return `(${str})`
}