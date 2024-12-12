package handlers

import (
	"api-crud/data"
	"api-crud/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Listar todos os produtos
func GetProducts(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data.Produtos); err != nil {
		http.Error(w, "Erro ao processar os dados", http.StatusInternalServerError)
		log.Printf("Erro ao codificar produtos: %v", err)
		return
	}
}

// Obter um produto por ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := getIDFromRequest(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	produto, found := findProductByID(id)
	if !found {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(produto); err != nil {
		http.Error(w, "Erro ao processar os dados", http.StatusInternalServerError)
		log.Printf("Erro ao codificar produto: %v", err)
	}
}

// Criar um novo produto
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	var newProduct models.Produto
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	// Gerar ID único
	newProduct.ID = getNextID()
	data.Produtos = append(data.Produtos, newProduct)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newProduct); err != nil {
		http.Error(w, "Erro ao processar os dados", http.StatusInternalServerError)
		log.Printf("Erro ao codificar novo produto: %v", err)
	}
}

// Atualizar um produto
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	id, err := getIDFromRequest(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, produto := range data.Produtos {
		if produto.ID == id {
			var updatedProduct models.Produto
			if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
				http.Error(w, "Requisição inválida", http.StatusBadRequest)
				return
			}

			// Atualizar o produto existente
			updatedProduct.ID = id
			data.Produtos[i] = updatedProduct

			if err := json.NewEncoder(w).Encode(updatedProduct); err != nil {
				http.Error(w, "Erro ao processar os dados", http.StatusInternalServerError)
				log.Printf("Erro ao codificar produto atualizado: %v", err)
			}
			return
		}
	}

	http.Error(w, "Produto não encontrado", http.StatusNotFound)
}

// Deletar um produto
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := getIDFromRequest(r)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, produto := range data.Produtos {
		if produto.ID == id {
			// Remover o produto
			data.Produtos = append(data.Produtos[:i], data.Produtos[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // Código 204: Sem conteúdo
			return
		}
	}

	http.Error(w, "Produto não encontrado", http.StatusNotFound)
}

// Função auxiliar para obter o ID da requisição
func getIDFromRequest(r *http.Request) (int, error) {
	params := mux.Vars(r)
	return strconv.Atoi(params["id"])
}

// Função auxiliar para buscar um produto por ID
func findProductByID(id int) (models.Produto, bool) {
	for _, produto := range data.Produtos {
		if produto.ID == id {
			return produto, true
		}
	}
	return models.Produto{}, false
}

// Gerar o próximo ID (simples incremento)
func getNextID() int {
	if len(data.Produtos) == 0 {
		return 1
	}
	return data.Produtos[len(data.Produtos)-1].ID + 1
}
