let gigamillisecond = 1e12

export const gigasecond = (date) => {
    return new Date(Number(date) + gigamillisecond)
};
