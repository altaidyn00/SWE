function num(value) {
  const regex = /^(0|[1-9][0-9]*)$/;
  return regex.test(value);
}

function positiveNum(value) {
  return value > 0;
}

function phoneNum(value) {
  const regex =
    /^\s*(?:\+?(\d{1,3}))?[-. (]*(\d{3})[-. )]*(\d{3})[-. ]*(\d{4})(?: *x(\d+))?\s*$/;
  return regex.test(value);
}

export { num, positiveNum, phoneNum };
