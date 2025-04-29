package repository

import (
	"TechShop/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TechOrderPostgres struct {
	db *sqlx.DB
}

func NewTechOrderPostgres(db *sqlx.DB) *TechOrderPostgres {
	return &TechOrderPostgres{db: db}
}
func (r *TechOrderPostgres) GetOrders(userId string) ([]models.Order, error) {
	var orders []models.Order
	query := fmt.Sprintf(`SELECT o.id, o.user_id, o.item_id, o.item_count, o.status,o.is_active,o.created_at, o.updated_at FROM %s o WHERE user_id=$1`, ordersTable)
	err := r.db.Select(&orders, query, userId)
	if err != nil {
		return nil, err
	}
	return orders, err
}
func (r *TechOrderPostgres) GetOrderById(userId, orderId string) (models.Order, error) {
	var order models.Order
	query := fmt.Sprintf(`SELECT o.id, o.user_id, o.item_id, o.item_count, o.status,o.is_active,o.created_at, o.updated_at FROM %s o WHERE o.user_id=$1 AND o.id=$2`, ordersTable)
	err := r.db.Get(&order, query, userId, orderId)
	if err != nil {
		return order, err
	}
	return order, err
}
func (r *TechOrderPostgres) CreateOrdersFromCart(userId string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	// 1. Забираем все товары из корзины
	var cartItems []models.CartItem // структура модели корзины
	getCartQuery := fmt.Sprintf(`SELECT id AS cart_id, user_id, item_id AS cart_item_id, item_count, created_at AS cart_created_at, updated_at AS cart_updated_at FROM %s WHERE user_id=$1`, cartTable)

	err = tx.Select(&cartItems, getCartQuery, userId)
	if err != nil {
		return err
	}
	if len(cartItems) == 0 {
		return fmt.Errorf("корзина пуста")
	}
	// 2. Вставляем в заказы
	insertOrderQuery := fmt.Sprintf(`
		INSERT INTO %s (user_id, item_id, item_count, status, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
	`, ordersTable)

	for _, cartItem := range cartItems {
		_, err := tx.Exec(insertOrderQuery,
			cartItem.UserID,
			cartItem.ItemID,
			cartItem.ItemCount,
			"pending", // статус по умолчанию, например
			true,      // активный заказ
		)
		if err != nil {
			return err
		}
	}

	// 3. Очищаем корзину
	deleteCartQuery := fmt.Sprintf(`DELETE FROM %s WHERE user_id=$1`, cartTable)
	_, err = tx.Exec(deleteCartQuery, userId)
	if err != nil {
		return err
	}

	return nil
}
func (r *TechOrderPostgres) CreateOrderFromCart(userId, cartId string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	// 1. Забираем все товары из корзины
	var cartItem models.CartItem // структура модели корзины
	getCartQuery := fmt.Sprintf(`SELECT id AS cart_id, user_id, item_id AS cart_item_id, item_count, created_at AS cart_created_at, updated_at AS cart_updated_at FROM %s WHERE user_id=$1 AND id=$2`, cartTable)

	err = tx.Get(&cartItem, getCartQuery, userId, cartId)
	if err != nil {
		return err
	}
	// 2. Вставляем в заказы
	insertOrderQuery := fmt.Sprintf(`
		INSERT INTO %s (user_id, item_id, item_count, status, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
	`, ordersTable)
	_, err = tx.Exec(insertOrderQuery,
		cartItem.UserID,
		cartItem.ItemID,
		cartItem.ItemCount,
		"pending", // статус по умолчанию, например
		true,      // активный заказ
	)
	if err != nil {
		return err
	}

	// 3. Очищаем корзину
	deleteCartQuery := fmt.Sprintf(`DELETE FROM %s WHERE user_id=$1 AND id=$2`, cartTable)
	_, err = tx.Exec(deleteCartQuery, userId, cartId)
	if err != nil {
		return err
	}

	return nil
}
func (r *TechOrderPostgres) CancelOrder(userId, orderId string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	query := fmt.Sprintf(`UPDATE %s SET status='canceled' WHERE user_id=$1 AND id=$2`, ordersTable)
	_, err = tx.Exec(query, userId, orderId)
	if err != nil {
		return err
	}
	queryIsActive := fmt.Sprintf(`UPDATE %s SET is_active=false WHERE user_id=$1 AND id=$2`, ordersTable)
	_, err = tx.Exec(queryIsActive, userId, orderId)
	if err != nil {
		return err
	}
	return nil
}
func (r *TechOrderPostgres) DeliveringOrder(userId, orderId string) error {
	query := fmt.Sprintf(`UPDATE %s SET status='delivering' WHERE user_id=$1 AND id=$2`, ordersTable)
	_, err := r.db.Exec(query, userId, orderId)
	if err != nil {
		return err
	}
	return nil
}
func (r *TechOrderPostgres) DeliveredOrder(userId, orderId string) error {
	query := fmt.Sprintf(`UPDATE %s SET status='delivered' WHERE user_id=$1 AND id=$2`, ordersTable)
	_, err := r.db.Exec(query, userId, orderId)
	if err != nil {
		return err
	}
	return nil
}
