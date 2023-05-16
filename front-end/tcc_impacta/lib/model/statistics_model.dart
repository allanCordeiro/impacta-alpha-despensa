import 'dart:convert';

class GetStats {
  Future <List<ProductList>?> getStatistics() async {
   var url = Uri.parse('https://despensa.onrender.com/api/stock/statistics');
    var http;
    var response = await http.get(url);
    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Não foi possível carregar a lista');
    }
  }
}


class GetListStatistcs {
  int? minimalQuantity;
  int? affectedProducts;
  List<ProductList>? productList;

  GetListStatistcs(
      {this.minimalQuantity, this.affectedProducts, this.productList});

  GetListStatistcs.fromJson(Map<String, dynamic> json) {
    minimalQuantity = json['minimal_quantity'];
    affectedProducts = json['affected_products'];
    if (json['product_list'] != null) {
      productList = <ProductList>[];
      json['product_list'].forEach((v) {
        productList!.add(ProductList.fromJson(v));
      });
    }
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['minimal_quantity'] = minimalQuantity;
    data['affected_products'] = affectedProducts;
    if (productList != null) {
      data['product_list'] = productList!.map((v) => v.toJson()).toList();
    }
    return data;
  }
}

class ProductList {
  String? id;
  String? name;

  ProductList({this.id, this.name});

  ProductList.fromJson(Map<String, dynamic> json) {
    id = json['id'];
    name = json['name'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['id'] = id;
    data['name'] = name;
    return data;
  }
}