import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

import '../../model/item_model.dart';

class HttpApiRepository with ChangeNotifier {
  final List<ItemModel> _items = [];
  final String _baseUrl = 'https://despensa.onrender.com/api/stock';
  
  Future<dynamic> getList() async {
    var url = Uri.parse(_baseUrl);
    var response = await http.get(url);
    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Não foi possível carregar a lista');
    }
  }

  Future<dynamic> getStatistics() async {
   var url = Uri.parse('https://despensa.onrender.com/api/stock/statistics');
    var response = await http.get(url);
    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Não foi possível carregar a lista');
    }
  }

  Future<void> addItem(ItemModel item) async {
    try {
      final response = await http.post(
        Uri.parse('https://despensa.onrender.com/api/stock'),
        body: json.encode(
          {
            "creation_date": item.creationDate,
            "expiration_date": item.expirationDate,
            "name": item.name,
            "quantity": item.quantity,
          },
        ),
      );
      final responseData = json.decode(response.body);
      final newItem = ItemModel.fromJson(responseData);
      _items.add(newItem);
      notifyListeners();
    } catch (e) {
      print(e);
    }
  }

  Future<void> removeItem(String id) async {
    try {
      var request = http.Request(
        'PUT',
        Uri.parse('https://despensa.onrender.com/api/products/$id/decrease'),
      );
      request.body = '''''';
      http.StreamedResponse response = await request.send();
      if (response.statusCode == 200) {
        print(await response.stream.bytesToString());
      } else {
        print(response.reasonPhrase);
      }
    } catch (e) {
      print(e);
    }
  }
}





// class ProductProvider with ChangeNotifier {
//   final List<ItemModel> _Items = [];

//   Future<dynamic> getList() async {
//     var url = Uri.parse(_baseUrl);
//     var response = await http.get(url);
//     if (response.statusCode == 200) {
//       return jsonDecode(response.body);
//     } else {
//       throw Exception('Não foi possível carregar a lista');
//     }
//   }

  // Future<void> fetchProducts() async {
  //   final response = await http.get(Uri.parse(_baseUrl));
  //   final List<dynamic> productsData = json.decode(response.body);
  //   final List<Product> loadedProducts = [];
  //   for (var productData in productsData) {
  //     loadedProducts.add(Product.fromJson(productData));
  //   }
  //   _products = loadedProducts;
  //   notifyListeners();
  // }

  // Future<void> addProduct(Product product) async {
  //   final response = await http.post(
  //     Uri.parse(_baseUrl),
  //     body: json.encode(product.toJson()),
  //     headers: {'Content-Type': 'application/json'},
  //   );
  //   final responseData = json.decode(response.body);
  //   final newProduct = Product.fromJson(responseData);
  //   _products.add(newProduct);
  //   notifyListeners();
  // }

  // Future<void> updateProduct(Product product) async {
  //   final productIndex = _products.indexWhere((p) => p.id == product.id);
  //   if (productIndex >= 0) {
  //     final response = await http.put(
  //       Uri.parse('$_baseUrl/${product.id}'),
  //       body: json.encode(product.toJson()),
  //       headers: {'Content-Type': 'application/json'},
  //     );
  //     _products[productIndex] = product;
  //     notifyListeners();
  //   }
  // }

  // Future<void> deleteProduct(int productId) async {
  //   final response = await http.delete(Uri.parse('$_baseUrl/$productId'));
  //   _products.removeWhere((p) => p.id == productId);
  //   notifyListeners();
  // }
// }