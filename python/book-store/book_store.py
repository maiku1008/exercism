from collections import Counter

DISCOUNTED = {
    1: 1 * 800 * 1.00,  # 800
    2: 2 * 800 * 0.95,  # 1520
    3: 3 * 800 * 0.90,  # 2160
    4: 4 * 800 * 0.80,  # 2560
    5: 5 * 800 * 0.75,  # 3000
}


def total(basket):
    return get_total(sorted(basket))


def get_total(books):
    basket = Counter(books)
    total = len(books) * 800
    for count in range(len(basket), 1, -1):
        group = basket - Counter(item for item, _ in basket.most_common(count))
        group_books = sorted(group.elements())
        total = min(total, DISCOUNTED[count] + get_total(group_books))
    return total
