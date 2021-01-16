from collections import Counter

SINGLE_ITEM_PRICE = 800
DISCOUNTED = {
    1: 1 * SINGLE_ITEM_PRICE * 1.00,  # 800
    2: 2 * SINGLE_ITEM_PRICE * 0.95,  # 1520
    3: 3 * SINGLE_ITEM_PRICE * 0.90,  # 2160
    4: 4 * SINGLE_ITEM_PRICE * 0.80,  # 2560
    5: 5 * SINGLE_ITEM_PRICE * 0.75,  # 3000
}


def total(basket: list) -> int:
    return get_total(basket)


def get_total(books: list) -> int:
    basket = Counter(books)
    total = len(books) * SINGLE_ITEM_PRICE
    for count in range(len(basket), 1, -1):
        group = basket - Counter(item for item, _ in basket.most_common(count))
        group_books = list(group.elements())
        total = min(total, DISCOUNTED[count] + get_total(group_books))
    return total
