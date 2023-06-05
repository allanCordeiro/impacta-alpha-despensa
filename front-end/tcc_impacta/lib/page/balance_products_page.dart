import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class BalanceProducts extends StatefulWidget {
  const BalanceProducts({Key? key}) : super(key: key);

  @override
  State<BalanceProducts> createState() => _BalanceProductsState();
}

class _BalanceProductsState extends State<BalanceProducts> {
  Future<List<dynamic>> getListBalances() async {
    var response = await http
        .get(Uri.parse('https://despensa.onrender.com/api/products/balance'));

    if (response.statusCode == 200) {
      dynamic data = jsonDecode(response.body);
      List<Map<String, dynamic>> result =
          []; // Movido para fora do loop forEach

      data.forEach((key, value) {
        print('Data: $key');
        value.forEach((item) {
          final name = item['name'];
          final deductedQuantity = item['deducted_quantity'];
          final remainingQuantity = item['remaining_quantity'];
          result.add({
            'date': key,
            'name': name,
            'deducted_quantity': deductedQuantity,
            'remaining_quantity': remainingQuantity,
          });
        });
      });

      return result;
    } else {
      throw Exception('Failed to load balance');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Histórico'),
      ),
      body: FutureBuilder<List<dynamic>>(
        future: getListBalances(),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          } else if (snapshot.hasError) {
            return Center(
              child: Text('Erro: ${snapshot.error}'),
            );
          } else if (snapshot.hasData) {
            List<dynamic> data = snapshot.data!;
            if (data.isNotEmpty) {
              // Agrupar dados por data
              Map<String, List<dynamic>> groupedData = {};
              for (dynamic item in data) {
                String date = item['date']
                    .toString(); // Supondo que a data está no campo 'date'
                if (groupedData.containsKey(date)) {
                  groupedData[date]!.add(item);
                } else {
                  groupedData[date] = [item];
                }
              }

              return Padding(
                padding: const EdgeInsets.all(8.0),
                child: ListView.builder(
                  itemCount: groupedData.length,
                  itemBuilder: (context, index) {
                    String date = groupedData.keys.elementAt(index);
                    List<dynamic> items = groupedData[date]!;

                    return Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        ExpansionTile(
                          title: Text(
                            date,
                            style: const TextStyle(
                              fontSize: 18,
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                          children: [
                            ListView.builder(
                              shrinkWrap: true,
                              physics: const NeverScrollableScrollPhysics(),
                              itemCount: items.length,
                              itemBuilder: (context, index) {
                                return Card(
                                  elevation: 3,
                                  child: Column(
                                    crossAxisAlignment:
                                        CrossAxisAlignment.center,
                                    children: [
                                      Row(
                                        children: [
                                          Text(
                                            '   Item: ${items[index]['name'].toString()}',
                                            style: const TextStyle(
                                              fontSize: 18,
                                              fontFamily: 'Futura',
                                            ),
                                          ),
                                        ],
                                      ),
                                      Row(
                                        children: [
                                          Text(
                                            '   Consumindo: ${items[index]['deducted_quantity'].toString()}',
                                            style: const TextStyle(
                                              fontSize: 18,
                                              fontFamily: 'Futura',
                                              color: Colors.red,
                                            ),
                                          ),
                                        ],
                                      ),
                                      Row(
                                        children: [
                                          Text(
                                            '   Estoque: ${items[index]['remaining_quantity'].toString()}',
                                            style: const TextStyle(
                                              fontSize: 18,
                                              fontFamily: 'Futura',
                                              color: Colors.brown,
                                            ),
                                          ),
                                        ],
                                      ),
                                    ],
                                  ),
                                );
                              },
                            ),
                          ],
                        ),
                      ],
                    );
                  },
                ),
              );
            } else {
              return const Center(
                child: Text('Nenhum dado disponível'),
              );
            }
          } else {
            return const Center(
              child: Text('Nenhum dado disponível'),
            );
          }
        },
      ),
    );
  }
}
