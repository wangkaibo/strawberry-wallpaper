// eslint-disable-next-line import/prefer-default-export
export const dateFormat = function (myDate, format) {
    const date = {
        'M+': myDate.getMonth() + 1,
        'd+': myDate.getDate(),
        'h+': myDate.getHours(),
        'm+': myDate.getMinutes(),
        's+': myDate.getSeconds(),
        'q+': Math.floor((myDate.getMonth() + 3) / 3),
        'S+': myDate.getMilliseconds(),
    };
    if (/(y+)/i.test(format)) {
        format = format.replace(RegExp.$1, (`${myDate.getFullYear()}`).substr(4 - RegExp.$1.length));
    }
    for (const k in date) {
        if (new RegExp(`(${k})`).test(format)) {
            format = format.replace(RegExp.$1, RegExp.$1.length === 1
                ? date[k] : (`00${date[k]}`).substr((`${date[k]}`).length));
        }
    }
    return format;
};
