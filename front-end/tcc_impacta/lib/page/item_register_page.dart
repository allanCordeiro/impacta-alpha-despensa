import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:tcc_impacta/model/item_model.dart';

import '../provider/item.dart';
import '../repository/implementations/http_api_repository.dart';
import '../routes/app_routes.dart';

class ItemRegister extends StatefulWidget {
  const ItemRegister({super.key});

  @override
  _ItemRegisterState createState() => _ItemRegisterState();
}

class _ItemRegisterState extends State<ItemRegister> {
  final List<ItemModel> _items = [];

  final _formKey = GlobalKey<FormState>();
  final Map<String, dynamic> _formData = {};
  @override
  Widget build(BuildContext context) {
    final HttpApiRepository httpPost = HttpApiRepository();
    final ItemProvider itemProvider = Provider.of(context);
    return Scaffold(
      appBar: AppBar(
        title: const Text('Cadastro'),
      ),
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: Card(
              color: Colors.teal[100],
              child: Padding(
                padding: const EdgeInsets.all(8.0),
                child: Form(
                  key: _formKey,
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      TextFormField(
                        onSaved: (value) => _formData['name'] = value!,
                        decoration: const InputDecoration(
                          labelText: 'Nome',
                        ),
                      ),
                      TextFormField(
                        onSaved: (value) =>
                            _formData['creationDate'] = value!,
                        decoration: const InputDecoration(
                          labelText: 'Data da compra',
                        ),
                      ),
                      TextFormField(
                        onSaved: (value) => _formData['quantity'] = value!,
                        decoration: const InputDecoration(
                          labelText: 'Quantidade',
                        ),
                        keyboardType: TextInputType.number,
                      ),
                      TextFormField(
                        onSaved: (value) => _formData['expirationDate'] = value!,
                        // initialValue: _expirationDate.toString(),
                        decoration: const InputDecoration(
                          labelText: 'Data de validade',
                        ),
                      ),
                      const SizedBox(height: 16),
                      Center(
                        child: ElevatedButton(
                          onPressed: () async{
                            _formKey.currentState?.save();
                            await httpPost.addItem(ItemModel(
                              name: _formData['name'] ?? '',
                              quantity: _formData['quantity'] ?? '',
                              creationDate: _formData['creationDate'] ?? '',
                              expirationDate: _formData['expirationDate'] ?? '',
                            ));
                            Navigator.of(context).pushReplacementNamed(AppRoutes.HOME);
                          },
                          child: const Text('Cadastrar'),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
