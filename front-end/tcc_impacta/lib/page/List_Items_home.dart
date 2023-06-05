import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:tcc_impacta/page/balance_iten_page.dart';
import 'package:tcc_impacta/page/statistics_list.dart';

import '../repository/implementations/http_api_repository.dart';
import '../routes/app_routes.dart';
import 'balance_products_page.dart';

class ListItems extends StatefulWidget {
  const ListItems({Key? key}) : super(key: key);

  @override
  _ListItemsState createState() => _ListItemsState();
}

class _ListItemsState extends State<ListItems> {
  int _notificationCount = 0;
  bool _isLoading = true;

  @override
  void initState() {
    super.initState();
    getStatistics();
  }

  Future<void> getStatistics() async {
    setState(() {
      _isLoading = true;
    });

    try {
      final response = await http
          .get(Uri.parse('https://despensa.onrender.com/api/stock/statistics'));
      if (response.statusCode == 200) {
        final jsonResponse = json.decode(response.body);
        final affectedProducts = jsonResponse['affected_products'];
        setState(() {
          _notificationCount = affectedProducts;
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(
              content: Text("So tem 1"),
              duration: Duration(seconds: 2),
            ),
          );
        });
      } else {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Falha ao carregar dados do servidor')),
        );
      }
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(
            content: Text('Ocorreu um erro ao obter as estatísticas')),
      );
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Lista de Itens'),
        leading: IconButton(
          icon: const Icon(Icons.history),
          onPressed: () {
            Navigator.push(
              context,
              MaterialPageRoute(
                builder: (context) =>  const BalanceProducts(),
              ),
            );
          },
        ),
        actions: [
          Stack(
            children: [
              IconButton(
                icon: const Icon(Icons.notifications),
                onPressed: () {
                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) => const ListStatistics(),
                    ),
                  );
                },
              ),
              Positioned(
                right: 4,
                child: Container(
                  padding: const EdgeInsets.all(2),
                  decoration: BoxDecoration(
                    color: Colors.red,
                    borderRadius: BorderRadius.circular(6),
                  ),
                  constraints: const BoxConstraints(
                    minWidth: 12,
                    minHeight: 12,
                  ),
                  child: Text(
                    '$_notificationCount',
                    style: const TextStyle(
                      color: Colors.white,
                      fontSize: 14,
                    ),
                    textAlign: TextAlign.center,
                  ),
                ),
              ),
            ],
          ),
        ],
      ),
      body: _isLoading
          ? const Center(child: CircularProgressIndicator())
          : const FutureListBuild(),
    );
  }
}

class FutureListBuild extends StatefulWidget {
  const FutureListBuild({
    Key? key,
  }) : super(key: key);

  @override
  State<FutureListBuild> createState() => _FutureListBuildState();
}

class _FutureListBuildState extends State<FutureListBuild> {
  final HttpApiRepository client = HttpApiRepository();
  @override
  Widget build(BuildContext context) {
    return FutureBuilder<dynamic>(
      future: client.getList(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          return Scaffold(
            body: GridView.builder(
              primary: false,
              padding: const EdgeInsets.all(20),
              itemCount: snapshot.data!.length,
              gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 2,
                childAspectRatio: 4 / 3.4,
              ),
              itemBuilder: (BuildContext context, int index) {
                return GestureDetector(
                  onTap: () async {
                    String id = snapshot.data[index]['id'];
                    Navigator.push(
                      context,
                      MaterialPageRoute(
                        builder: (context) =>
                            BalanceItemPage(id: id), // Passar o ID aqui
                      ),
                    );
                  },
                  child: Card(
                    elevation: 5,
                    child: Column(
                      children: [
                        Padding(
                          padding: const EdgeInsets.all(8.0),
                          child: Text(
                            snapshot.data[index]['name'],
                            style: const TextStyle(fontSize: 20),
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.all(8.0),
                          child: Text(
                            snapshot.data[index]['expiration_date']
                                .split('-')
                                .reversed
                                .join('-'),
                            style: const TextStyle(fontSize: 16),
                          ),
                        ),
                        Row(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: [
                            Row(
                              children: [
                                Padding(
                                  padding: const EdgeInsets.symmetric(
                                      horizontal: 6.0),
                                  child: Text(snapshot.data[index]['quantity'],
                                      style: const TextStyle(fontSize: 18)),
                                ),
                              ],
                            ),
                            IconButton(
                              onPressed: () async {
                                await client
                                    .removeItem(snapshot.data[index]['id']);
                                setState(() {});
                                ScaffoldMessenger.of(context).showSnackBar(
                                  const SnackBar(
                                    content: Text("Item excluído com sucesso!"),
                                    duration: Duration(seconds: 2),
                                  ),
                                );
                              },
                              icon: const Icon(Icons.delete),
                              color: Colors.red,
                            )
                          ],
                        ),
                      ],
                    ),
                  ),
                );
              },
            ),
            floatingActionButton: FloatingActionButton.small(
              onPressed: () {
                Navigator.of(context).pushNamed(AppRoutes.ITEMS_FORM);
              },
              child: const Icon(Icons.fact_check_outlined),
            ),
          );
        } else if (snapshot.hasError) {
          return Center(child: Text('${snapshot.error}'));
        }
        return const Center(child: CircularProgressIndicator());
      },
    );
  }
}
