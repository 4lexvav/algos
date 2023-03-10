package asteriod

func AsteroidCollision(asteroids []int) []int {
	for i := 0; i < len(asteroids); i++ {
		weight := asteroids[i]
		if weight > 0 || i == 0 {
			continue
		}

	explodeAsteroids:
		for j := i - 1; j >= 0; j-- {
			// exit, there no positive weight asteroids left
			if asteroids[j] < 0 {
				break
			}

			delta := asteroids[j] + weight
			switch true {
			// asteroids of same size: explode both and exit loop
			case delta == 0:
				asteroids = append(asteroids[:j], asteroids[i+1:]...)
				i -= 2
				break explodeAsteroids
			// positive asteroid is greater: explode only negative and exit loop
			case delta > 0:
				if i+1 == len(asteroids) {
					asteroids = asteroids[:i]
				} else {
					asteroids = append(asteroids[:i], asteroids[i+1:]...)
					i--
				}
				break explodeAsteroids
			// negative asteroid is greater: explode only positive and try to explode next asteroid
			case delta < 0:
				asteroids = append(asteroids[:j], asteroids[i:]...)
				i--
			}
		}
	}

	return asteroids
}
