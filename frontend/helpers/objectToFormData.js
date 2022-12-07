function createFormData(obj, formData, subKeyStr = "") {
  for (const i in obj) {
    const value = obj[i];
    const subKeyStrTrans = subKeyStr ? subKeyStr + "[" + i + "]" : i;
    if (
      typeof value === "string" ||
      typeof value === "number" ||
      value instanceof File
    ) {
      formData.append(subKeyStrTrans, value);
    } else if (typeof value === "object") {
      createFormData(value, formData, subKeyStrTrans);
    }
  }
}
export default function objectToFormData(obj, formData = new FormData()) {
  createFormData(obj, formData);
  return formData;
}
