const EARTH_SECONDS_IN_YEAR = 315576 * Math.pow(10, 2);
const EARTH_YEARS_ON_PLANET = {
  mercury: 0.2408467,
  venus: 0.61519726,
  earth: 1.0,
  mars: 1.8808158,
  jupiter: 11.862615,
  saturn: 29.447498,
  uranus: 84.016846,
  neptune: 164.79132,
}

export const age = (planet, seconds) => {
  return round(seconds / EARTH_SECONDS_IN_YEAR / EARTH_YEARS_ON_PLANET[planet], 2)
}

const round = (value, decimals) => {
  const precision = Math.pow(10, decimals);
  const epsilon = Number.EPSILON;
  return Math.round((value + epsilon) * precision) / precision;
}
