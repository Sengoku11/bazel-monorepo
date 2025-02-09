import numpy as np
import scipy.stats

def analyze_data(x: np.ndarray, y: np.ndarray):
    regression = scipy.stats.linregress(x, y)

    return {
        "slope": regression.slope,
        "intercept": regression.intercept,
        "r_squared": regression.rvalue ** 2,
    }

if __name__ == "__main__":
    x = np.array([1, 2, 3, 4, 5])
    y = np.array([2, 4, 5, 4, 5])

    results = analyze_data(x, y)
    print(results)